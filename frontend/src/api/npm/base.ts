import { camelizeKeys, decamelizeKeys } from "humps";
import AuthStore from "modules/AuthStore";
import * as queryString from "query-string";

interface BuildUrlArgs {
	url: string;
	params?: queryString.StringifiableRecord;
}

function buildUrl({ url, params }: BuildUrlArgs) {
	const endpoint = url.replace(/^\/|\/$/g, "");
	const apiParams = params ? `?${queryString.stringify(params)}` : "";
	const apiUrl = `/api/${endpoint}${apiParams}`;
	return apiUrl;
}

function buildAuthHeader(): Record<string, string> | undefined {
	if (AuthStore.token) {
		return { Authorization: `Bearer ${AuthStore.token.token}` };
	}
	return {};
}

function buildBody(data?: Record<string, any>) {
	if (data) {
		return JSON.stringify(decamelizeKeys(data));
	}
}

async function processResponse(response: Response) {
	const payload = await response.json();
	if (!response.ok) {
		throw new Error(payload.error.message);
	}
	return camelizeKeys(payload) as any;
}

interface GetArgs {
	url: string;
	params?: queryString.StringifiableRecord;
}

export async function get(
	{ url, params }: GetArgs,
	abortController?: AbortController,
) {
	const apiUrl = buildUrl({ url, params });
	const method = "GET";
	const signal = abortController?.signal;
	const headers = buildAuthHeader();
	const response = await fetch(apiUrl, { method, headers, signal });
	return processResponse(response);
}

interface PostArgs {
	url: string;
	data?: any;
}

export async function post(
	{ url, data }: PostArgs,
	abortController?: AbortController,
) {
	const apiUrl = buildUrl({ url });
	const method = "POST";
	const headers = { ...buildAuthHeader(), "Content-Type": "application/json" };
	const signal = abortController?.signal;
	const body = buildBody(data);
	const response = await fetch(apiUrl, { method, headers, body, signal });
	return processResponse(response);
}

interface PutArgs {
	url: string;
	data?: any;
}

export async function put(
	{ url, data }: PutArgs,
	abortController?: AbortController,
) {
	const apiUrl = buildUrl({ url });
	const method = "PUT";
	const headers = { ...buildAuthHeader(), "Content-Type": "application/json" };
	const signal = abortController?.signal;
	const body = buildBody(data);
	const response = await fetch(apiUrl, { method, headers, body, signal });
	return processResponse(response);
}
