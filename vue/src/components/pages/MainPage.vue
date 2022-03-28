<template>
    <main-page
        :tampering-message="tampering"
        :check-date="lastCheckDate"
        :notarizations="notarizationData"
        :data="statusData">
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
            return this.statusData.some(item => item.status !== 'NORMAL')
                ? 'Tampering Detected'
                : 'No Tampering Detected';
        },

        lastCheckDate() {
            const lastCheckTime = this.statusData[this.statusData.length - 1]?.time;
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

            if (data[hash].length > 30) {
                const slicedArray =  data[hash].slice(-30);
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

           data[hash].forEach((item, index) => {
               item.key = index;
               // item.group = 'Dataset 1';
               item.collectTime = formattedDateLocaleString(item.collectTime);
               // item.newNotarizationsCount = numFormatter(item.newNotarizationsCount);
           });

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
