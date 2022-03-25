<template>
    <main-page
        :tampering-message="tampering"
        :check-date="checkDate"
        :data="data">
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
            data: [],
            checkDate: ''
        };
    },

    beforeMount() {
        this.data = [
            {
                "signer_id": "",
                "time": "2022-03-23T23:54:12.557569+03:00",
                "status": "NORMAL",
                "prev_tx_id": 0,
                "prev_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9",
                "new_tx_id": 0,
                "new_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9"
            },
            {
                "signer_id": "",
                "time": "2022-03-23T23:55:12.556986+03:00",
                "status": "NORMAL",
                "prev_tx_id": 0,
                "prev_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9",
                "new_tx_id": 0,
                "new_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9"
            },
            {
                "signer_id": "",
                "time": "2022-03-23T23:56:12.558209+03:00",
                "status": "NORMAL",
                "prev_tx_id": 0,
                "prev_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9",
                "new_tx_id": 0,
                "new_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9"
            },
            {
                "signer_id": "",
                "time": "2022-03-23T23:57:12.5566+03:00",
                "status": "NORMAL",
                "prev_tx_id": 0,
                "prev_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9",
                "new_tx_id": 0,
                "new_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9"
            },
            {
                "signer_id": "",
                "time": "2022-03-23T23:58:12.559218+03:00",
                "status": "NORMAL",
                "prev_tx_id": 0,
                "prev_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9",
                "new_tx_id": 0,
                "new_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9"
            },
            {
                "signer_id": "",
                "time": "2022-03-23T23:59:12.55631+03:00",
                "status": "NORMAL",
                "prev_tx_id": 0,
                "prev_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9",
                "new_tx_id": 0,
                "new_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9"
            },
            {
                "signer_id": "",
                "time": "2022-03-24T00:00:12.556658+03:00",
                "status": "NORMAL",
                "prev_tx_id": 0,
                "prev_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9",
                "new_tx_id": 0,
                "new_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9"
            },
            {
                "signer_id": "",
                "time": "2022-03-24T00:01:12.56043+03:00",
                "status": "NORMAL",
                "prev_tx_id": 0,
                "prev_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9",
                "new_tx_id": 0,
                "new_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9"
            },
            {
                "signer_id": "",
                "time": "2022-03-24T00:02:12.559673+03:00",
                "status": "NORMAL",
                "prev_tx_id": 0,
                "prev_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9",
                "new_tx_id": 0,
                "new_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9"
            },
            {
                "signer_id": "",
                "time": "2022-03-24T00:03:12.556593+03:00",
                "status": "NORMAL",
                "prev_tx_id": 0,
                "prev_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9",
                "new_tx_id": 0,
                "new_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9"
            },
            {
                "signer_id": "",
                "time": "2022-03-24T00:04:12.561582+03:00",
                "status": "NORMAL",
                "prev_tx_id": 0,
                "prev_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9",
                "new_tx_id": 0,
                "new_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9"
            },
            {
                "signer_id": "",
                "time": "2022-03-24T00:05:12.573969+03:00",
                "status": "NORMAL",
                "prev_tx_id": 0,
                "prev_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9",
                "new_tx_id": 0,
                "new_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9"
            },
            {
                "signer_id": "",
                "time": "2022-03-24T00:06:12.580635+03:00",
                "status": "NORMAL",
                "prev_tx_id": 0,
                "prev_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9",
                "new_tx_id": 0,
                "new_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9"
            },
            {
                "signer_id": "",
                "time": "2022-03-24T00:07:12.577694+03:00",
                "status": "NORMAL",
                "prev_tx_id": 0,
                "prev_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9",
                "new_tx_id": 0,
                "new_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9"
            },
            {
                "signer_id": "",
                "time": "2022-03-24T00:08:12.5783+03:00",
                "status": "NORMAL",
                "prev_tx_id": 0,
                "prev_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9",
                "new_tx_id": 0,
                "new_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9"
            },
            {
                "signer_id": "",
                "time": "2022-03-24T00:09:12.583235+03:00",
                "status": "NORMAL",
                "prev_tx_id": 0,
                "prev_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9",
                "new_tx_id": 0,
                "new_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9"
            },
            {
                "signer_id": "",
                "time": "2022-03-24T11:27:59.530615+03:00",
                "status": "NORMAL",
                "prev_tx_id": 0,
                "prev_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9",
                "new_tx_id": 0,
                "new_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9"
            },
            {
                "signer_id": "",
                "time": "2022-03-24T11:28:59.526262+03:00",
                "status": "NORMAL",
                "prev_tx_id": 0,
                "prev_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9",
                "new_tx_id": 0,
                "new_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9"
            },
            {
                "signer_id": "",
                "time": "2022-03-24T11:29:59.52612+03:00",
                "status": "NORMAL",
                "prev_tx_id": 0,
                "prev_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9",
                "new_tx_id": 0,
                "new_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9"
            },
            {
                "signer_id": "",
                "time": "2022-03-24T11:30:59.52802+03:00",
                "status": "NORMAL",
                "prev_tx_id": 0,
                "prev_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9",
                "new_tx_id": 0,
                "new_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9"
            },
            {
                "signer_id": "",
                "time": "2022-03-24T11:31:59.525126+03:00",
                "status": "NORMAL",
                "prev_tx_id": 0,
                "prev_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9",
                "new_tx_id": 0,
                "new_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9"
            },
            {
                "signer_id": "",
                "time": "2022-03-24T11:32:59.526745+03:00",
                "status": "NORMAL",
                "prev_tx_id": 0,
                "prev_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9",
                "new_tx_id": 0,
                "new_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9"
            },
            {
                "signer_id": "",
                "time": "2022-03-24T11:33:59.526766+03:00",
                "status": "NORMAL",
                "prev_tx_id": 0,
                "prev_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9",
                "new_tx_id": 0,
                "new_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9"
            },
            {
                "signer_id": "",
                "time": "2022-03-24T11:34:59.529344+03:00",
                "status": "NORMAL",
                "prev_tx_id": 0,
                "prev_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9",
                "new_tx_id": 0,
                "new_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9"
            },
            {
                "signer_id": "",
                "time": "2022-03-24T11:35:59.526668+03:00",
                "status": "NORMAL",
                "prev_tx_id": 0,
                "prev_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9",
                "new_tx_id": 0,
                "new_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9"
            },
            {
                "signer_id": "",
                "time": "2022-03-24T11:36:59.528913+03:00",
                "status": "NORMAL",
                "prev_tx_id": 0,
                "prev_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9",
                "new_tx_id": 0,
                "new_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9"
            },
            {
                "signer_id": "",
                "time": "2022-03-24T11:37:59.52563+03:00",
                "status": "NORMAL",
                "prev_tx_id": 0,
                "prev_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9",
                "new_tx_id": 0,
                "new_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9"
            },
            {
                "signer_id": "",
                "time": "2022-03-24T11:38:59.526605+03:00",
                "status": "NORMAL",
                "prev_tx_id": 0,
                "prev_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9",
                "new_tx_id": 0,
                "new_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9"
            },
            {
                "signer_id": "",
                "time": "2022-03-24T11:39:59.547573+03:00",
                "status": "NORMAL",
                "prev_tx_id": 0,
                "prev_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9",
                "new_tx_id": 0,
                "new_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9"
            },
            {
                "signer_id": "",
                "time": "2022-03-24T11:40:59.546699+03:00",
                "status": "NORMAL",
                "prev_tx_id": 0,
                "prev_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9",
                "new_tx_id": 0,
                "new_state_hash": "0ff36aea5ee2194f2f025092a7bc3fe978c10d217c7dee5497f742a5b4468fd9"
            }
        ]

        this.getLastCheckDate();

        // fetch('http://localhost:8091/api/status')
        //     .then(response => response.json())
        //     .then(data => console.log(data));
        // this.$axios.get('http://localhost:8091/api/status').then(data => console.log(data))
    },

    computed: {
        tampering() {
            return this.data.some(item => item.status !== 'NORMAL')
                ? 'Tampering Detected'
                : 'No Tampering Detected';
        }
    },

    methods: {
        getLastCheckDate() {
            const lastCheckTime = this.data[this.data.length - 1].time;
            const date = new Date(lastCheckTime);

            this.checkDate = `${date.toDateString()} at ${date.toTimeString().split(' ')[0]}`;
        }
    }
}
</script>
