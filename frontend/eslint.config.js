import js from '@eslint/js'
import globals from 'globals'
import reactHooks from 'eslint-plugin-react-hooks'
import reactRefresh from 'eslint-plugin-react-refresh'
import tseslint from 'typescript-eslint'
import {defineConfig, globalIgnores} from 'eslint/config'

export default defineConfig([
    globalIgnores(['dist']),
    {
        files: ['**/*.{ts,tsx}'],
        extends: [
            js.configs.recommended,
            tseslint.configs.recommended,
            reactHooks.configs.flat.recommended,
            reactRefresh.configs.vite,
        ],
        languageOptions: {
            ecmaVersion: 2020,
            globals: globals.browser,
        },
        rules: {
            // Enforce spacing inside curly braces
            'object-curly-spacing': ['error', 'always'],

            // Optional but commonly paired rules
            'array-bracket-spacing': ['error', 'never'],
            'computed-property-spacing': ['error', 'never'],
            'block-spacing': ['error', 'always'],
            indent: ['error', 2],
            'no-mixed-spaces-and-tabs': 'error',
        },
    },
])
