<template>
    <main-page
        :tampering-message="tampering"
        :check-date="lastCheckDate"
        :data="statusData">
    </main-page>
</template>

<script>
import MainPage from '../templates/MainPage';

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
            return this.statusData.some(item => item.status !== 'NORMAL')
                ? 'Tampering Detected'
                : 'No Tampering Detected';
        },

        lastCheckDate() {
            const lastCheckTime = this.statusData[this.statusData.length - 1].time;
            const date = new Date(lastCheckTime);

            return `${date.toDateString()} at ${date.toTimeString().split(' ')[0]}`;
        }
    },

    methods: {
        async fetchStatus() {
           const { data } = await this.$axios.get('http://localhost:8091/api/status');

            if (!data) {
                return;
            }

            const hash = Object.keys(data)[0];
            this.statusData = data[hash];

            console.log('STATUS DATA:', this.statusData);
        },

        async fetchNotarizationCount() {
           const { data } = await this.$axios.get('http://localhost:8091/api/notarization/count');

           if (!data) {
               return;
           }

            const hash = Object.keys(data)[0];
            this.notarizationData = data[hash];

            console.log('NOTARIZATION DATA:', this.notarizationData);
        }
    }
}
</script>
