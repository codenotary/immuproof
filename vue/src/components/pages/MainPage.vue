<template>
    <main-page
        :tampering-message="tampering"
        :check-date="lastCheckDate"
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

        lastCheckDate() {
            const lastCheckTime = this.statusData[this.statusData.length - 1]?.time;
            const date = new Date(lastCheckTime);

            return `${date.toDateString()} at ${date.toTimeString().split(' ')[0]}`;
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
                const slicedArray =  data[hash].slice(-45);
                this.statusData = slicedArray;

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
                const slicedArray =  data[hash].slice(-30);
                this.notarizationData = slicedArray;

                return;
            }

           this.notarizationData = data[hash];
        }
    }
}
</script>
