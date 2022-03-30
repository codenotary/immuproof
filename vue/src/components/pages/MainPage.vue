<template>
    <main-page
        :tampering-message="tampering"
        :last-check-date="lastCheckDate"
        :last-tx-id="lastTXId"
        :first-check-date="firstCheckDate"
        :notarizations="notarizationData"
        :data="statusData"
        :notarization-count-categories="notarizationCountCategories"
        :notarization-count-data="notarizationCountData">
    </main-page>
</template>

<script>
import MainPage from '../templates/MainPage';
import { formattedDateLocaleString } from '@/helpers/helpers';

export default {
    components: {
        MainPage
    },

    data() {
        return {
            statusData: [],
            notarizationData: []
        };
    },

    async beforeMount() {
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
        notarizationCountCategories() {
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
            const prefix = process.env.NODE_ENV === 'development' ? 'http://localhost:8091/' : '/';
            const { data } = await this.$axios.get(`${prefix}api/notarization/count`);

            if (!data) {
                return;
            }

            const hash = Object.keys(data)[0];

            if (data[hash].length > 45) {
                this.statusData = data[hash].slice(-45);

                return;
            }

            this.statusData = data[hash];
        },

        async fetchNotarizationCount() {
            const prefix = process.env.NODE_ENV === 'development' ? 'http://localhost:8091/' : '/';
            const { data } = await this.$axios.get(`${prefix}api/notarization/count`);

            if (!data) {
                return;
            }

            const hash = Object.keys(data)[0];

            if (data[hash].length > 30) {
                this.notarizationData = data[hash].slice(-30);

                return;
            }

            this.notarizationData = data[hash];
        }
    }
}
</script>
