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
                    zoom: {
                        enabled: false
                    },
                    toolbar: {
                        show: false
                    }
                },
                stroke: {
                    curve: 'smooth'
                },
                xaxis: {
                    categories: this.categories,
                    labels: {
                        show: false
                    }
                },
                yaxis: {
                    labels: {
                        show: false
                    }
                },
                dataLabels: {
                    enabled: false
                }
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

<style scoped>

</style>