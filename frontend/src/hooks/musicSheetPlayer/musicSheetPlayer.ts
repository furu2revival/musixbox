import type { MusicSheet } from "~/model/musicSheet";
import type { Pitch } from "~/model/note";
import {
	createPooledAudioPlayer,
	PooledAudioPlayer,
} from "./pooledAudioPlayer";

import soundFileA4 from "~/assets/A4.mp3";
import soundFileB4 from "~/assets/B4.mp3";
import soundFileC3 from "~/assets/C3.mp3";
import soundFileC4 from "~/assets/C4.mp3";
import soundFileD3 from "~/assets/D3.mp3";
import soundFileE3 from "~/assets/E3.mp3";
import soundFileF3 from "~/assets/F3.mp3";
import soundFileG3 from "~/assets/G3.mp3";

const soundFiles: Record<Pitch, string> = {
	C3: soundFileC3,
	D3: soundFileD3,
	E3: soundFileE3,
	F3: soundFileF3,
	G3: soundFileG3,
	A4: soundFileA4,
	B4: soundFileB4,
	C4: soundFileC4,
};

export type MusicSheetPlayerInit = {
	musicSheet: MusicSheet;
	beatsPerMinute: number;
	maxEnergy: number;
	audioPlayerPerNode: number;
};

export class MusicSheetPlayer extends EventTarget {
	private audioPlayers: Record<Pitch, PooledAudioPlayer>;
	private musicSheet: MusicSheet;
	/** Ability to play music (unit: beats) */
	private _energy = 0;
	private maxEnergy: number;
	private noteInterval: number;

	private indexInSheet = 0;
	private timeout: number | undefined = undefined;

	get energy() {
		return this._energy;
	}

	set energy(value: number) {
		this._energy = value;
		this.dispatchEvent(new Event("energychange"));
	}

	constructor(init: MusicSheetPlayerInit) {
		super();

		this.musicSheet = init.musicSheet;
		this.maxEnergy = init.maxEnergy;
		this.noteInterval = (60 / init.beatsPerMinute) * 1000;

		this.audioPlayers = {} as typeof this.audioPlayers;
		this.load(init.audioPlayerPerNode);
	}

	private async load(poolSize: number) {
		await Promise.all(
			Object.entries(soundFiles).map(async ([pitch, file]) => {
				this.audioPlayers[pitch as Pitch] = await createPooledAudioPlayer(
					file,
					poolSize
				);
			})
		);
		this.dispatchEvent(new Event("load"));
	}

	play() {
		this.timeout = setTimeout(() => this.playNext(), 0);
	}

	stop() {
		clearTimeout(this.timeout);
		this.timeout = undefined;
	}

	setEnergy(energy: number) {
		this.energy = Math.min(energy, this.maxEnergy);
	}

	private playNext() {
		if (this.energy < 1) {
			this.energy = 0;
			this.stop();
			return;
		}
		this.energy -= 1;

		const note = this.musicSheet.notes[this.indexInSheet];
		console.log(note);
		for (const pitch of note.pitch) {
			this.audioPlayers[pitch].play();
		}

		this.indexInSheet = (this.indexInSheet + 1) % this.musicSheet.notes.length;
		this.timeout = setTimeout(() => this.playNext(), this.noteInterval);
	}
}

export const createMusicSheetPlayer = async function (
	init: MusicSheetPlayerInit
): Promise<MusicSheetPlayer> {
	return new Promise((resolve) => {
		const player = new MusicSheetPlayer(init);
		player.addEventListener("load", () => {
			resolve(player);
		});
	});
};