<template>
    <main-page
        :tampering-message="tampering"
        :last-check-date="lastCheckDate"
        :last-tx-id="lastTXId"
        :first-check-date="firstCheckDate"
        :notarizations="notarizationData"
        :data="statusData"
        :logo-url="logoUrl"
        :notarization-categories-count="notarizationCategoriesCount"
        :notarization-count-data="notarizationCountData">
    </main-page>
</template>

<script>
import { formattedDateLocaleString } from '@/helpers/helpers';
import MainPage from '@/components/templates/MainPage.vue';

const MAX_STATUS_NUMBER = 45;
const MAX_NOTARIZATION_NUMBER = 30;

export default {
    components: {
        MainPage
    },
    data() {
        return {
            statusData: [],
            notarizationData: [],
            logoUrl: '',
            portValue: '',
            address: ''
        };
    },
    beforeCreate() {
        document.title = 'Immuproof';
    },
    async beforeMount() {
        this.checkLogoUrl();

        await Promise.all([
            this.fetchStatus(),
            this.fetchNotarizationCount()
        ]);
    },
    computed: {
        tampering() {
            if (this.statusData.some(item => item.status === 'CORRUPTED_DATA')) {
                return 'Tampering Detected';
            }

            return this.statusData[this.statusData.length - 1]?.status === 'NORMAL'
                ? 'No Tampering Detected'
                : 'Status Unknown';
        },
        firstCheckDate() {
            const firstCheckTime = this.statusData[0]?.time;

            return formattedDateLocaleString(firstCheckTime,
                {
                    year: 'numeric',
                    weekday: 'long',
                    month: 'long',
                    day: '2-digit',
                    hour: 'numeric',
                    minute: 'numeric',
                    timeZoneName: 'short'
                });
        },
        lastData() {
            return this.statusData[this.statusData.length - 1];
        },
        lastCheckDate() {
            const lastCheckTime = this.lastData?.time;

            return formattedDateLocaleString(lastCheckTime,
                {
                    year: 'numeric',
                    weekday: 'long',
                    month: 'long',
                    day: '2-digit',
                    hour: 'numeric',
                    minute: 'numeric',
                    timeZoneName: 'short'
                });
        },
        lastTXId() {
            return this.lastData?.new_tx_id;
        },
        notarizationCategoriesCount() {
            return this.notarizationData.map(data =>
                formattedDateLocaleString(
                    data.collectTime,
                    { month: 'long', day: 'numeric', hour: '2-digit', minute: '2-digit' }
                )
            );
        },
        notarizationCountData() {
            return this.notarizationData.map(data => data.newNotarizationsCount);
        }
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
            if (hostedByLogoUrl && hostedByLogoUrl.includes('{{')) {
                this.logoUrl = '';

                return;
            }

            this.logoUrl = hostedByLogoUrl;
        }
    }
}
</script>
