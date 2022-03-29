<template>
    <div class="hover-box">
        <div class="hover-box__wrapper">
            <div class="hover-box__wrapper-title">
                {{ title }}
            </div>
            <div class="hover-box__wrapper-time">
                {{ formattedTime }}
            </div>
            <div class="hover-box__wrapper-information">
                <div class="hover-box__wrapper-subtitle">
                    Previous TxID
                </div>
                <div class="hover-box__wrapper-content">
                    {{ historyData.prev_tx_id }}
                </div>
                <div class="hover-box__wrapper-subtitle">
                    Previous State Hash
                </div>
                <div class="hover-box__wrapper-content">
                    {{ historyData.prev_state_hash }}
                </div>
                <div class="hover-box__wrapper-subtitle">
                    New TxID
                </div>
                <div class="hover-box__wrapper-content">
                    {{ historyData.new_tx_id }}
                </div>
                <div class="hover-box__wrapper-subtitle">
                    New State Hash
                </div>
                <div class="hover-box__wrapper-content">
                    {{ historyData.new_state_hash }}
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { formattedDateLocaleString } from '@/helpers/helpers';

export default {
    props: {
        title: {
            type: String,
            default: ''
        },
        historyData: {
            type: Object,
            default: () => ({})
        }
    },
    computed: {
        formattedTime() {
            return formattedDateLocaleString(this.historyData.time, {
                year: 'numeric',
                weekday: 'long',
                month: 'long',
                day: '2-digit',
                hour: 'numeric',
                minute: 'numeric',
                timeZoneName: 'short'
            });
        }
    }
}
</script>