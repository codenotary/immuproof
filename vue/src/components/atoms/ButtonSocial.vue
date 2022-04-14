<template>
	<div class="buttons-social d-flex">
		<v-btn
			:icon="true"
			href="https://github.com/codenotary/cas"
			rel="noopener"
			target="_blank"
			title="Visit the cas repo on GitHub"
			:color="color"
		>
			<v-img
				contain
				:src="require('@/assets/images/icon-github.svg')"
				width="24px"
			>
			</v-img>
		</v-btn>
	</div>
</template>

<script>
export default {
	props: {
		color: {
			type: String,
			default: 'white',
		},
	},
};
</script>
