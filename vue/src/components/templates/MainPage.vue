<template>
    <div class="main-page">
        <floating-menu></floating-menu>
        <v-container>
            <v-row class="text-center">
                <cas-header></cas-header>
                <div class="main-page__content full-width">
                    <tampering-status
                        :tampering-message="tamperingMessage"
                        :check-date="checkDate">
                    </tampering-status>
                    <tampering-history
                        :check-date="checkDate"
                        :data="data">
                    </tampering-history>
                    <notarizations :notarizations="notarizations"></notarizations>
                    <apex-chart type="bar"></apex-chart>
                </div>
            </v-row>
        </v-container>
        <cas-footer></cas-footer>
    </div>
</template>

<script>
import FloatingMenu from '@/components/organisms/FloatingMenu';
import CasHeader from '@/components/organisms/CasHeader';
import TamperingStatus from '@/components/organisms/TamperingStatus';
import TamperingHistory from '@/components/organisms/TamperingHistory';
import Notarizations from '@/components/organisms/Notarizations';
import CasFooter from '@/components/organisms/CasFooter';
import ApexChart from "@/components/organisms/ApexChart";

export default {
    components: {
        ApexChart,
        CasFooter,
        Notarizations,
        TamperingStatus,
        CasHeader,
        FloatingMenu,
        TamperingHistory
    },

    props: {
        tamperingMessage: {
            type: String,
            required: true
        },
        checkDate: {
            type: String,
            required: true
        },
        data: {
            type: Array,
            default: () => ([])
        },
        notarizations: {
            type: Array,
            default: () => ([])
        }
    },
}
</script>