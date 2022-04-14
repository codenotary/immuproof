<template>
	<div class="cas-header elevation-6 text-center full-width pa-8">
		<h2 class="text-uppercase font-weight-bold">{{ titleText }}</h2>
		<v-img
			class="cas-header__stamp"
			:src="require('@/assets/images/img-stamp.svg')"
			height="auto"
			:contain="true"
		/>
		<div class="cas-header__hosted-by d-flex justify-center align-center mt-4">
			<p class="cas-header__hosted-by-text ma-0 mr-2 font-weight-bold">{{ hostedByText }}</p>
			<a
				:href="logoLink"
				target="_blank"
			>
				<v-img
					v-if="showInternalLogo"
					:src="require('@/assets/images/logo-only-small.svg')"
					:max-height="30"
					:max-width="50"
					:contain="true"
					height="auto"
				>
				</v-img>
				<v-img
					v-else
					:src="logoUrl"
					:max-height="30"
					:max-width="50"
					:contain="true"
					height="auto"
				>
				</v-img>
			</a>
		</div>
	</div>
</template>

<script>
export default {
	props: {
		logoUrl: { type: String, default: '' },
		logoLink: { type: String, default: '' },
		hostedByText: { type: String, default: 'Hosted by:' },
		titleText: { type: String, default: 'Community Attestation Service Status Page' },
	},

	computed: {
		showInternalLogo() {
			return this.logoUrl === '';
		},
	},
};
</script>

<style scoped lang="scss">
    .cas-header {
        border-radius: 0 0 24px 24px;

        h2 {
            line-height: 1;
            color: #244583;
        }
    }
</style>