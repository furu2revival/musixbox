import { useQuery } from "@tanstack/react-query";
import { getV1 } from "~/sdk/api/music_sheet-MusicSheetService_connectquery";

export const useMusicSheet = (id: string) => {
	const { isLoading, isError, error, data } = useQuery(
		getV1.useQuery({
			musicSheetId: id,
		}),
	);
	return {
		isLoading,
		isError,
		error,
		data,
	};
};