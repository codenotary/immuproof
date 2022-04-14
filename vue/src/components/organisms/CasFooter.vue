<template>
	<div class="cas-footer mt-8">
		<div class="cas-footer__logos pa-3 d-flex justify-space-between">
			<div class="cas-footer__main-logo">
				<v-img
					:src="require('@/assets/images/logo.svg')"
					height="auto"
					max-width="172px"
					:contain="true"
				>
				</v-img>
			</div>
			<div class="cas-footer__socials d-flex">
				<button-social></button-social>
				<button-secured class="ml-3"></button-secured>
			</div>
		</div>
		<div class="cas-footer__copyright text-center py-3">
			Copyright © 2022 Codenotary, Inc. All rights reserved.
		</div>
	</div>
</template>

<script>
import ButtonSocial from '@/components/atoms/ButtonSocial.vue';
import ButtonSecured from '@/components/atoms/ButtonSecured.vue';

export default {
	components: {
		ButtonSecured,
		ButtonSocial,
	},
};
</script>

<style scoped>

</style>