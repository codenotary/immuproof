<template>
    <ccv-line-chart
        :data="data"
        :options="options">
    </ccv-line-chart>
</template>

<script>
export default {
    name: "Chart",
    data() {
        return {
            data: [
                {
                    group: "Dataset 1",
                    key: 0,
                    value: 47263
                },
                {
                    group: "Dataset 1",
                    key: 1,
                    value: 44178
                },
                {
                    group: "Dataset 1",
                    key: 2,
                    value: 23094
                },
                {
                    group: "Dataset 1",
                    key: 3,
                    value: 45281
                },
                {
                    group: "Dataset 1",
                    key: 4,
                    value: 63954
                }
            ],
            options: {
                legend: {
                    enabled: false
                },
                axes: {
                    bottom: {
                        mapsTo: "key",
                        scaleType: "time",
                    },
                    left: {
                        mapsTo: "value",
                        scaleType: "linear"
                    },
                },
                "height": "170px",
                "toolbar": {
                    enabled: false
                }
            }
        };
    }
}
</script>

<style scoped>

</style>