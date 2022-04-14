<template>
	<v-btn
		class="main-button d-flex align-center justify-start overflow-hidden mt-2"
		target="_blank"
		rel="noopener"
		:href="link"
		:depressed="depressed"
		:outlined="outlined"
		:color="outlined ? color : ''"
	>
		<slot
			v-if="icon !== ''"
			name="icon"
		>
			<v-icon
				:style="{ color: color }"
				class="mr-2"
			>
				{{ icon }}
			</v-icon>
		</slot>
		<v-divider
			v-if="divider"
			class="main-button__divider my-4"
			:color="color"
			:vertical="true"
		>
		</v-divider>
		<slot></slot>
	</v-btn>
</template>

<script>
export default {
	props: {
		link: { type: String, default: '' },
		outlined: { type: Boolean, default: true },
		depressed: { type: Boolean, default: true },
		divider: { type: Boolean, default: false },
		color: { type: String, default: '#ffffff' },
		icon: { type: String, default: '' },
	},
};
</script>