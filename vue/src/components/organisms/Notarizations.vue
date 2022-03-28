<template>
    <div class="notarizations full-width mt-10">
        <v-container class="notarizations-wrapper">
            <v-row>
                <div class="col-12 notarizations-wrapper__texts">
                    <div class="text-left font-weight-bold">Notarizations</div>
                    <div class="text-left">Since Monday, 12 January 2022 at 15:34</div>
                </div>
                <apex-chart
                    class="mt-3"
                    type="line"
                    :categories="notarizationCountCategories"
                    :data="notarizationCountData">
                </apex-chart>
            </v-row>
        </v-container>
    </div>
</template>

<script>
import ApexChart from '@/components/organisms/ApexChart';

export default {
    components: {
        ApexChart
    },
    props: {
        notarizations: {
            type: Array,
            default: () => ([])
        },
        notarizationCountCategories: {
            type: Array,
            default: () => ([])
        },
        notarizationCountData: {
            type: Array,
            default: () => ([])
        }
    },
    computed: {
        options() {
            return {
                legend: { enabled: false },
                axes: {
                    bottom: { title: 'Time', scaleType: 'labels', visible: true },
                    left: { title: 'Count', mapsTo: 'newNotarizationsCount', scaleType: 'linear' },
                },
                height: '170px',
                toolbar: { enabled: false }
            }
        }
    }
}
</script>