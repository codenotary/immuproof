<template>
    <div class="histogram-line d-flex justify-space-between align-center full-width col-12">
        <div>{{ firstCheckDateFormatted }}</div>
        <span class="histogram-line__line"></span>
        <div>{{ lastCheckDateFormatted }}</div>
    </div>
</template>

<script>
import { formattedDateLocaleString } from '@/helpers/helpers';

export default {
    props: {
        lastCheckDate: { type: String, required: true },
        firstCheckDate: { type: String, required: true }
    },
    computed: {
        lastCheckDateFormatted() {
            return formattedDateLocaleString(this.lastCheckDate, {
                month: 'short',
                timeZoneName: 'short'
            });
        },
        firstCheckDateFormatted() {
            return formattedDateLocaleString(this.firstCheckDate, {
                month: 'short',
                timeZoneName: 'short'
            });
        }
    }
}
</script>

<style scoped>

</style>