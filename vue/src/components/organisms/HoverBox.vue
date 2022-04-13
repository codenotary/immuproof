<template>
    <div class="hover-box">
        <div class="hover-box__title">
            {{ title }}
        </div>
        <div class="hover-box__time">
            {{ formattedTime }}
        </div>
        <div class="hover-box__information">
            <div class="hover-box__subtitle">
                Previous Transaction ID
            </div>
            <div class="hover-box__content">
                {{ historyData.prev_tx_id }}
            </div>
            <div class="hover-box__subtitle">
                Previous State Hash
            </div>
            <div class="hover-box__content">
                {{ historyData.prev_state_hash }}
            </div>
            <div class="hover-box__subtitle">
                New Transaction ID
            </div>
            <div class="hover-box__content">
                {{ historyData.new_tx_id }}
            </div>
            <div class="hover-box__subtitle">
                New State Hash
            </div>
            <div class="hover-box__content">
                {{ historyData.new_state_hash }}
            </div>
        </div>
    </div>
</template>

<script>
import { formattedDateLocaleString } from '@/helpers/helpers';

export default {
    props: {
        title: { type: String, default: '' },
        historyData: { type: Object, default: () => ({}) }
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