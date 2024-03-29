<template>
	<main-page
		:tampering-message="tampering"
		:last-check-date="lastCheckDate"
		:last-tx-id="lastTXId"
		:utc-check-first="utcCheckFirst"
		:utc-check-last="utcCheckLast"
		:notarizations="notarizationData"
		:data="statusData"
		:logo-url="logoUrl"
		:logo-link="logoLink"
		:hosted-by-text="hostedByText"
		:title-text="titleText"
		:notarization-categories-count="notarizationCategoriesCount"
		:notarization-count-data="notarizationCountData"
	>
	</main-page>
</template>

<script>
import { formattedDateLocaleString } from '@/helpers/helpers';
import MainPage from '@/components/templates/MainPage.vue';

const MAX_STATUS_NUMBER = 45;
const MAX_NOTARIZATION_NUMBER = 30;

export default {
	components: {
		MainPage,
	},
	data() {
		return {
			statusData: [],
			notarizationData: [],
			logoUrl: '',
			logoLink: '',
			hostedByText: '',
			titleText: '',
			portValue: '',
			address: '',
		};
	},
	beforeCreate() {
		document.title = 'CAS Validator';
	},
	async beforeMount() {
		this.checkLogoUrl();
		this.checkHostedByText();
		this.checkTitleText();
		this.checkHostedByLogoLink();

		await Promise.all([
			this.fetchStatus(),
			this.fetchNotarizationCount(),
		]);
	},
	computed: {
		tampering() {
			if (this.statusData.some(item => item.status === 'CORRUPTED_DATA')) {
				return 'Validation not successful';
			}

			return this.statusData[this.statusData.length - 1]?.status === 'NORMAL'
				? 'Validation successful'
				: 'Validation not successful';
		},
		utcCheckFirst() {
			return this.statusData[0]?.time;
		},
		lastData() {
			return this.statusData[this.statusData.length - 1];
		},
		utcCheckLast() {
			return this.lastData?.time;
		},
		lastCheckDate() {
			return formattedDateLocaleString(this.utcCheckLast,
				{
					year: 'numeric',
					weekday: 'long',
					month: 'long',
					day: '2-digit',
					hour: 'numeric',
					minute: 'numeric',
					timeZoneName: 'short',
				});
		},
		lastTXId() {
			return this.lastData?.new_tx_id;
		},
		notarizationCategoriesCount() {
			return this.notarizationData.map(data =>
				formattedDateLocaleString(
					data.collectTime,
					{ month: 'long', day: 'numeric', hour: '2-digit', minute: '2-digit' },
				),
			);
		},
		notarizationCountData() {
			return this.notarizationData.map(data => data.newNotarizationsCount);
		},
	},
	methods: {
		async fetchStatus() {
			const prefix = this.getAddressPrefix();
			const { data } = await this.$axios.get(`${prefix}/api/status`);

			if (!data) {
				return;
			}

			const hash = Object.keys(data)[0];

			if (data[hash].length > MAX_STATUS_NUMBER) {
				this.statusData = data[hash].slice(-MAX_STATUS_NUMBER);

				return;
			}

			this.statusData = data[hash];
		},
		async fetchNotarizationCount() {
			const prefix = this.getAddressPrefix();
			const { data } = await this.$axios.get(`${prefix}/api/notarization/count`);

			if (!data) {
				return;
			}

			const hash = Object.keys(data)[0];

			if (data[hash].length > MAX_NOTARIZATION_NUMBER) {
				this.notarizationData = data[hash].slice(-MAX_NOTARIZATION_NUMBER);

				return;
			}

			this.notarizationData = data[hash];
		},
		getAddressPrefix() {
			if (process.env.NODE_ENV === 'development') {
				const { PORT = '8091' } = process.env;

				return `http://localhost:${PORT}`;
			}

			return '';
		},
		checkLogoUrl() {
			this.logoUrl = hostedByLogoUrl?.includes('{{') ? '' : hostedByLogoUrl;
		},
		checkHostedByText() {
			this.hostedByText = hostedByText?.includes('{{') ? 'Hosted by:' : hostedByText;
		},
		checkTitleText() {
			this.titleText = titleText?.includes('{{') ? 'Community Attestation Service Validator' : titleText;
		},
		checkHostedByLogoLink() {
			this.logoLink = hostedByLogoLink?.includes('{{') ?'https://cas.codenotary.com/' : hostedByLogoLink;
		},
	},
};
</script>
