<template>
	<div class="tampering-status full-width mt-10">
		<v-container>
			<v-row class="tampering-status__content align-center">
				<div
					class="tampering-status__status d-flex align-center"
					:class="getClassName"
				>
					<p class="ma-0 ml-3">{{ tamperingMessage }}</p>
				</div>
				<div class="tampering-status__last-check">
					<p class="font-weight-bold text-left">Last Check</p>
					<p class="tampering-status__last-check-date text-left">{{ lastCheckDate }}</p>
					<p class="tampering-status__last-check-id text-left">
						<span>Current Transaction ID</span>
						{{ lastTxId }}
					</p>
				</div>
			</v-row>
		</v-container>
	</div>
</template>

<script>
export default {
	props: {
		tamperingMessage: { type: String, required: true },
		lastCheckDate: { type: String, required: true },
		lastTxId: { type: Number, required: true },
	},

	computed: {
		getClassName() {
			switch (this.tamperingMessage) {
				case 'Status Unknown':
					return 'tampering-status__status-unknown';
				case 'Tampering Detected':
					return 'tampering-status__status-error';
				default:
					return 'tampering-status__status-normal';
			}
		},
	},
};
</script>

<style scoped>

</style>