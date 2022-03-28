<template>
    <main-page
        :tampering-message="tampering"
        :last-check-date="lastCheckDate"
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
        lastCheckDate() {
            const lastCheckTime = this.statusData[this.statusData.length - 1]?.time;

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
           const { data } = await this.$axios.get('http://localhost:8091/api/status');

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
           const { data } = await this.$axios.get('http://localhost:8091/api/notarization/count');

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
