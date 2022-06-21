import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'
import { viteSingleFile } from 'vite-plugin-singlefile'

export default defineConfig({
    plugins: [
        svelte(),
        viteSingleFile({
            removeViteModuleLoader: true,
            useRecommendedBuildConfig: true,
        }),
    ],
    build: {
        outDir: 'dist',
        target: 'es2019',
    },
    assetsInclude: ['**/*.png'],
})
