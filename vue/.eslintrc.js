module.exports = {
	root: true,
	env: {
		node: true,
	},
	'extends': [
		'plugin:vue/essential',
		'eslint:recommended',
	],
	globals: {
		'address': true,
		'portValue': true,
		'hostedByLogoUrl': true,
		'hostedByText': true,
		'titleText': true,
		'hostedByLogoLink': true,
	},
	parserOptions: {
		parser: '@babel/eslint-parser',
	},
	rules: {
		'no-console': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
		'no-debugger': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
		quotes: [
			'error',
			'single',
		],
		semi: [
			'error',
			'always',
		],
		indent: [
			'warn',
			'tab',
			{
				SwitchCase: 1,
				MemberExpression: 2,
			},
		],
		'array-element-newline': [
			'error', {
				ArrayExpression: 'consistent',
				ArrayPattern: { minItems: 1 },
			},
		],
		'array-bracket-newline': ['error', { multiline: true }],
		'space-before-function-paren': ['error', 'never'],
		'no-tabs': 'off',
		'vue/multi-word-component-names': 0,
		'comma-dangle': [
			'warn',
			{
				arrays: 'always-multiline',
				objects: 'always-multiline',
				imports: 'always-multiline',
				exports: 'always-multiline',
				functions: 'always-multiline',
			},
		],
		'vue/html-closing-bracket-newline': [
			'warn',
			{
				multiline: 'always',
				singleline: 'never',
			},
		],
		'vue/html-indent': [
			'warn',
			'tab',
			{
				baseIndent: 1,
				attribute: 1,
			},
		],
		'spaced-comment': 'warn',
		'vue/no-unused-vars': 'warn',
		'no-trailing-spaces': 'warn',
		'no-multiple-empty-lines': 'warn',
	},
};
