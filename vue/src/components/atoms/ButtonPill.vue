<template>
	<v-btn
		:href="href"
		class="button-pill | d-flex align-center white--light overflow-hidden ma-0 pa-0"
		target="_blank"
		rel="noopener"
		depressed
		:outlined="outlined"
		:color="outlined ? color : ''"
	>
		<slot name="icon">
			<v-icon
				:size="22"
				:class="`button-pill__icon ${ color }--text py-1 px-2`"
				:color="color"
			>
				{{ icon }}
			</v-icon>
		</slot>
		<v-divider
			v-if="divider"
			class="my-4"
			style="max-height: 20px; height: 20px;"
			:color="color"
			:vertical="true"
		/>
		<slot />
	</v-btn>
</template>

<script>
export default {
	name: 'ButtonPill',
	props: {
		href: { type: String, default: '' },
		icon: { type: String, default: '' },
		color: { type: String, default: 'primary' },
		outlined: { type: Boolean, default: false },
		divider: { type: Boolean, default: true },
	},
};
</script>

<style lang="scss">
.button-pill {
    .v-btn__content {
        justify-content: start;
    }

    &__icon {
        border-top-right-radius: 0;
        border-bottom-right-radius: 0;
    }

    .v-icon {
        margin-right: 0;
    }
}
</style>