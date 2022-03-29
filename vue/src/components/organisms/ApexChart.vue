<template>
    <div class="full-width">
        <apexchart
            :type="type"
            :options="chartOptions"
            :series="series">
        </apexchart>
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

    computed: {
        chartOptions() {
            return {
                chart: {
                    zoom: { enabled: false },
                    toolbar: { show: false },
                    height: '170px'
                },
                stroke: { curve: 'smooth', width: 2 },
                tooltip: { enabled: false },
                colors: ['#153954'],
                xaxis: { categories: this.categories, labels: { show: false } },
                yaxis: { labels: { show: false } },
                dataLabels: { enabled: false }
            };
        },
        series() {
            return [{
                name: 'Notarizations',
                data: this.data
            }]
        }
    }
};
</script>