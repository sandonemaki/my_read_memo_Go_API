import { defineConfig } from "orval";

export default defineConfig({
	yondeco: {
		input: {
			target: "../openapi/openapi.yml",
		},
		output: {
			mode: "split",
			target: "./app/api/endpoint",
			schemas: "./app/api/models",
			client: "swr",
			mock: true,
			clean: true,
		},
		hooks: {
			afterAllFilesWrite: "bun run --bun format",
		},
	},
});
