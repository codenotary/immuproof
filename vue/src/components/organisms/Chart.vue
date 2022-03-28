<template>
    <ccv-line-chart
        :data="data"
        :options="options">
    </ccv-line-chart>
</template>

<script>
export default {
    props: {
        data: {
            type: Array,
            required: true
        },
        options: {
            type: Object,
            required: true
        }
    }
}
</script>