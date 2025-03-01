// @generated by protoc-gen-es v1.10.0 with parameter "target=ts"
// @generated from file config/config.proto (package config, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type {
	BinaryReadOptions,
	FieldList,
	JsonReadOptions,
	JsonValue,
	PartialMessage,
	PlainMessage,
} from "@bufbuild/protobuf";
import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";

/**
 * @generated from message config.Config
 */
export class Config extends Message<Config> {
	/**
	 * @generated from field: bool debug = 1;
	 */
	debug = false;

	/**
	 * @generated from field: config.GoogleCloud google_cloud = 2;
	 */
	googleCloud?: GoogleCloud;

	/**
	 * @generated from field: config.Logging logging = 3;
	 */
	logging?: Logging;

	/**
	 * @generated from field: config.Postgres postgres = 4;
	 */
	postgres?: Postgres;

	/**
	 * @generated from field: config.APIServer api_server = 5;
	 */
	apiServer?: APIServer;

	constructor(data?: PartialMessage<Config>) {
		super();
		proto3.util.initPartial(data, this);
	}

	static readonly runtime: typeof proto3 = proto3;
	static readonly typeName = "config.Config";
	static readonly fields: FieldList = proto3.util.newFieldList(() => [
		{ no: 1, name: "debug", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
		{ no: 2, name: "google_cloud", kind: "message", T: GoogleCloud },
		{ no: 3, name: "logging", kind: "message", T: Logging },
		{ no: 4, name: "postgres", kind: "message", T: Postgres },
		{ no: 5, name: "api_server", kind: "message", T: APIServer },
	]);

	static fromBinary(
		bytes: Uint8Array,
		options?: Partial<BinaryReadOptions>
	): Config {
		return new Config().fromBinary(bytes, options);
	}

	static fromJson(
		jsonValue: JsonValue,
		options?: Partial<JsonReadOptions>
	): Config {
		return new Config().fromJson(jsonValue, options);
	}

	static fromJsonString(
		jsonString: string,
		options?: Partial<JsonReadOptions>
	): Config {
		return new Config().fromJsonString(jsonString, options);
	}

	static equals(
		a: Config | PlainMessage<Config> | undefined,
		b: Config | PlainMessage<Config> | undefined
	): boolean {
		return proto3.util.equals(Config, a, b);
	}
}

/**
 * @generated from message config.GoogleCloud
 */
export class GoogleCloud extends Message<GoogleCloud> {
	/**
	 * @generated from field: string project_id = 1;
	 */
	projectId = "";

	/**
	 * @generated from field: config.GoogleCloud.Trace trace = 2;
	 */
	trace?: GoogleCloud_Trace;

	constructor(data?: PartialMessage<GoogleCloud>) {
		super();
		proto3.util.initPartial(data, this);
	}

	static readonly runtime: typeof proto3 = proto3;
	static readonly typeName = "config.GoogleCloud";
	static readonly fields: FieldList = proto3.util.newFieldList(() => [
		{ no: 1, name: "project_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
		{ no: 2, name: "trace", kind: "message", T: GoogleCloud_Trace },
	]);

	static fromBinary(
		bytes: Uint8Array,
		options?: Partial<BinaryReadOptions>
	): GoogleCloud {
		return new GoogleCloud().fromBinary(bytes, options);
	}

	static fromJson(
		jsonValue: JsonValue,
		options?: Partial<JsonReadOptions>
	): GoogleCloud {
		return new GoogleCloud().fromJson(jsonValue, options);
	}

	static fromJsonString(
		jsonString: string,
		options?: Partial<JsonReadOptions>
	): GoogleCloud {
		return new GoogleCloud().fromJsonString(jsonString, options);
	}

	static equals(
		a: GoogleCloud | PlainMessage<GoogleCloud> | undefined,
		b: GoogleCloud | PlainMessage<GoogleCloud> | undefined
	): boolean {
		return proto3.util.equals(GoogleCloud, a, b);
	}
}

/**
 * @generated from message config.GoogleCloud.Trace
 */
export class GoogleCloud_Trace extends Message<GoogleCloud_Trace> {
	/**
	 * @generated from field: bool enabled = 1;
	 */
	enabled = false;

	/**
	 * @generated from field: float sampling_rate = 2;
	 */
	samplingRate = 0;

	constructor(data?: PartialMessage<GoogleCloud_Trace>) {
		super();
		proto3.util.initPartial(data, this);
	}

	static readonly runtime: typeof proto3 = proto3;
	static readonly typeName = "config.GoogleCloud.Trace";
	static readonly fields: FieldList = proto3.util.newFieldList(() => [
		{ no: 1, name: "enabled", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
		{
			no: 2,
			name: "sampling_rate",
			kind: "scalar",
			T: 2 /* ScalarType.FLOAT */,
		},
	]);

	static fromBinary(
		bytes: Uint8Array,
		options?: Partial<BinaryReadOptions>
	): GoogleCloud_Trace {
		return new GoogleCloud_Trace().fromBinary(bytes, options);
	}

	static fromJson(
		jsonValue: JsonValue,
		options?: Partial<JsonReadOptions>
	): GoogleCloud_Trace {
		return new GoogleCloud_Trace().fromJson(jsonValue, options);
	}

	static fromJsonString(
		jsonString: string,
		options?: Partial<JsonReadOptions>
	): GoogleCloud_Trace {
		return new GoogleCloud_Trace().fromJsonString(jsonString, options);
	}

	static equals(
		a: GoogleCloud_Trace | PlainMessage<GoogleCloud_Trace> | undefined,
		b: GoogleCloud_Trace | PlainMessage<GoogleCloud_Trace> | undefined
	): boolean {
		return proto3.util.equals(GoogleCloud_Trace, a, b);
	}
}

/**
 * @generated from message config.Logging
 */
export class Logging extends Message<Logging> {
	/**
	 * @generated from field: config.Logging.Severity severity = 1;
	 */
	severity = Logging_Severity.UNSPECIFIED;

	constructor(data?: PartialMessage<Logging>) {
		super();
		proto3.util.initPartial(data, this);
	}

	static readonly runtime: typeof proto3 = proto3;
	static readonly typeName = "config.Logging";
	static readonly fields: FieldList = proto3.util.newFieldList(() => [
		{
			no: 1,
			name: "severity",
			kind: "enum",
			T: proto3.getEnumType(Logging_Severity),
		},
	]);

	static fromBinary(
		bytes: Uint8Array,
		options?: Partial<BinaryReadOptions>
	): Logging {
		return new Logging().fromBinary(bytes, options);
	}

	static fromJson(
		jsonValue: JsonValue,
		options?: Partial<JsonReadOptions>
	): Logging {
		return new Logging().fromJson(jsonValue, options);
	}

	static fromJsonString(
		jsonString: string,
		options?: Partial<JsonReadOptions>
	): Logging {
		return new Logging().fromJsonString(jsonString, options);
	}

	static equals(
		a: Logging | PlainMessage<Logging> | undefined,
		b: Logging | PlainMessage<Logging> | undefined
	): boolean {
		return proto3.util.equals(Logging, a, b);
	}
}

/**
 * Cloud Logging に定義された LogSeverity に対応する。
 * https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry#logseverity
 *
 * @generated from enum config.Logging.Severity
 */
export enum Logging_Severity {
	/**
	 * @generated from enum value: SEVERITY_UNSPECIFIED = 0;
	 */
	UNSPECIFIED = 0,

	/**
	 * @generated from enum value: SEVERITY_DEBUG = 100;
	 */
	DEBUG = 100,

	/**
	 * @generated from enum value: SEVERITY_INFO = 200;
	 */
	INFO = 200,

	/**
	 * @generated from enum value: SEVERITY_NOTICE = 300;
	 */
	NOTICE = 300,

	/**
	 * @generated from enum value: SEVERITY_WARNING = 400;
	 */
	WARNING = 400,

	/**
	 * @generated from enum value: SEVERITY_ERROR = 500;
	 */
	ERROR = 500,

	/**
	 * @generated from enum value: SEVERITY_CRITICAL = 600;
	 */
	CRITICAL = 600,

	/**
	 * @generated from enum value: SEVERITY_ALERT = 700;
	 */
	ALERT = 700,

	/**
	 * @generated from enum value: SEVERITY_EMERGENCY = 800;
	 */
	EMERGENCY = 800,
}
// Retrieve enum metadata with: proto3.getEnumType(Logging_Severity)
proto3.util.setEnumType(Logging_Severity, "config.Logging.Severity", [
	{ no: 0, name: "SEVERITY_UNSPECIFIED" },
	{ no: 100, name: "SEVERITY_DEBUG" },
	{ no: 200, name: "SEVERITY_INFO" },
	{ no: 300, name: "SEVERITY_NOTICE" },
	{ no: 400, name: "SEVERITY_WARNING" },
	{ no: 500, name: "SEVERITY_ERROR" },
	{ no: 600, name: "SEVERITY_CRITICAL" },
	{ no: 700, name: "SEVERITY_ALERT" },
	{ no: 800, name: "SEVERITY_EMERGENCY" },
]);

/**
 * @generated from message config.Postgres
 */
export class Postgres extends Message<Postgres> {
	/**
	 * @generated from field: string host = 1;
	 */
	host = "";

	/**
	 * @generated from field: int32 port = 2;
	 */
	port = 0;

	/**
	 * @generated from field: string user = 3;
	 */
	user = "";

	/**
	 * @generated from field: string password = 4;
	 */
	password = "";

	/**
	 * @generated from field: string database = 5;
	 */
	database = "";

	/**
	 * @generated from field: string sslmode = 6;
	 */
	sslmode = "";

	constructor(data?: PartialMessage<Postgres>) {
		super();
		proto3.util.initPartial(data, this);
	}

	static readonly runtime: typeof proto3 = proto3;
	static readonly typeName = "config.Postgres";
	static readonly fields: FieldList = proto3.util.newFieldList(() => [
		{ no: 1, name: "host", kind: "scalar", T: 9 /* ScalarType.STRING */ },
		{ no: 2, name: "port", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
		{ no: 3, name: "user", kind: "scalar", T: 9 /* ScalarType.STRING */ },
		{ no: 4, name: "password", kind: "scalar", T: 9 /* ScalarType.STRING */ },
		{ no: 5, name: "database", kind: "scalar", T: 9 /* ScalarType.STRING */ },
		{ no: 6, name: "sslmode", kind: "scalar", T: 9 /* ScalarType.STRING */ },
	]);

	static fromBinary(
		bytes: Uint8Array,
		options?: Partial<BinaryReadOptions>
	): Postgres {
		return new Postgres().fromBinary(bytes, options);
	}

	static fromJson(
		jsonValue: JsonValue,
		options?: Partial<JsonReadOptions>
	): Postgres {
		return new Postgres().fromJson(jsonValue, options);
	}

	static fromJsonString(
		jsonString: string,
		options?: Partial<JsonReadOptions>
	): Postgres {
		return new Postgres().fromJsonString(jsonString, options);
	}

	static equals(
		a: Postgres | PlainMessage<Postgres> | undefined,
		b: Postgres | PlainMessage<Postgres> | undefined
	): boolean {
		return proto3.util.equals(Postgres, a, b);
	}
}

/**
 * @generated from message config.APIServer
 */
export class APIServer extends Message<APIServer> {
	/**
	 * @generated from field: int32 port = 1;
	 */
	port = 0;

	/**
	 * @generated from field: config.APIServer.Cors cors = 2;
	 */
	cors?: APIServer_Cors;

	constructor(data?: PartialMessage<APIServer>) {
		super();
		proto3.util.initPartial(data, this);
	}

	static readonly runtime: typeof proto3 = proto3;
	static readonly typeName = "config.APIServer";
	static readonly fields: FieldList = proto3.util.newFieldList(() => [
		{ no: 1, name: "port", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
		{ no: 2, name: "cors", kind: "message", T: APIServer_Cors },
	]);

	static fromBinary(
		bytes: Uint8Array,
		options?: Partial<BinaryReadOptions>
	): APIServer {
		return new APIServer().fromBinary(bytes, options);
	}

	static fromJson(
		jsonValue: JsonValue,
		options?: Partial<JsonReadOptions>
	): APIServer {
		return new APIServer().fromJson(jsonValue, options);
	}

	static fromJsonString(
		jsonString: string,
		options?: Partial<JsonReadOptions>
	): APIServer {
		return new APIServer().fromJsonString(jsonString, options);
	}

	static equals(
		a: APIServer | PlainMessage<APIServer> | undefined,
		b: APIServer | PlainMessage<APIServer> | undefined
	): boolean {
		return proto3.util.equals(APIServer, a, b);
	}
}

/**
 * @generated from message config.APIServer.Cors
 */
export class APIServer_Cors extends Message<APIServer_Cors> {
	/**
	 * @generated from field: repeated string allowed_origins = 1;
	 */
	allowedOrigins: string[] = [];

	/**
	 * @generated from field: repeated string allowed_methods = 2;
	 */
	allowedMethods: string[] = [];

	/**
	 * @generated from field: repeated string allowed_headers = 3;
	 */
	allowedHeaders: string[] = [];

	/**
	 * @generated from field: repeated string expose_headers = 4;
	 */
	exposeHeaders: string[] = [];

	/**
	 * @generated from field: int64 max_age = 5;
	 */
	maxAge = protoInt64.zero;

	/**
	 * @generated from field: bool allow_credentials = 6;
	 */
	allowCredentials = false;

	constructor(data?: PartialMessage<APIServer_Cors>) {
		super();
		proto3.util.initPartial(data, this);
	}

	static readonly runtime: typeof proto3 = proto3;
	static readonly typeName = "config.APIServer.Cors";
	static readonly fields: FieldList = proto3.util.newFieldList(() => [
		{
			no: 1,
			name: "allowed_origins",
			kind: "scalar",
			T: 9 /* ScalarType.STRING */,
			repeated: true,
		},
		{
			no: 2,
			name: "allowed_methods",
			kind: "scalar",
			T: 9 /* ScalarType.STRING */,
			repeated: true,
		},
		{
			no: 3,
			name: "allowed_headers",
			kind: "scalar",
			T: 9 /* ScalarType.STRING */,
			repeated: true,
		},
		{
			no: 4,
			name: "expose_headers",
			kind: "scalar",
			T: 9 /* ScalarType.STRING */,
			repeated: true,
		},
		{ no: 5, name: "max_age", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
		{
			no: 6,
			name: "allow_credentials",
			kind: "scalar",
			T: 8 /* ScalarType.BOOL */,
		},
	]);

	static fromBinary(
		bytes: Uint8Array,
		options?: Partial<BinaryReadOptions>
	): APIServer_Cors {
		return new APIServer_Cors().fromBinary(bytes, options);
	}

	static fromJson(
		jsonValue: JsonValue,
		options?: Partial<JsonReadOptions>
	): APIServer_Cors {
		return new APIServer_Cors().fromJson(jsonValue, options);
	}

	static fromJsonString(
		jsonString: string,
		options?: Partial<JsonReadOptions>
	): APIServer_Cors {
		return new APIServer_Cors().fromJsonString(jsonString, options);
	}

	static equals(
		a: APIServer_Cors | PlainMessage<APIServer_Cors> | undefined,
		b: APIServer_Cors | PlainMessage<APIServer_Cors> | undefined
	): boolean {
		return proto3.util.equals(APIServer_Cors, a, b);
	}
}
