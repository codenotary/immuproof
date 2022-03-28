<template>
    <div>
        <apexchart
            :type="type"
            :options="chartOptions"
            :series="series"></apexchart>
    </div>
</template>

<script>
export default {
    props: {
        type: {
            type: String,
            required: true
        },
        categories: {
            type: Array,
            default: () => ([])
        },
        data: {
            type: Array,
            default: () => ([])
        }
    },

    data() {
        return {
            chartOptions: {
                chart: {
                    id: 'vuechart-example'
                },
                xaxis: {
                    categories: [1991, 1992, 1993, 1994, 1995, 1996, 1997, 1998]
                }
            },
            series: [{
                name: 'series-1',
                data: [30, 40, 35, 50, 49, 60, 70, 91]
            }]
        }
    },
};
</script>

<style scoped>

</style>