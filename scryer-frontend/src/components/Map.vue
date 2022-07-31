<script setup>
import { storeToRefs } from 'pinia'

import LoginForm from './LoginForm.vue'
import RegistrationForm from './RegistrationForm.vue'

import { useDeviceStore } from '@/stores/device'
import { useDevicePreferencesStore } from '@/stores/devicePreferences'

const { devices, visibleDevices } =  storeToRefs(useDeviceStore())
</script>

<template>
	<LoginForm />
	<RegistrationForm />
	<GMapMap
		:center="center"
		:zoom="6"
	>
		<GMapMarker
			:key="device.device_id + '-marker'"
			v-for="device in visibleDevices"
			:position="device.latest_accurate_device_point.device_point_detail.lat_lng"
			@click="$emit('set-center', device.latest_accurate_device_point.device_point_detail.lat_lng)"
			:label="{ className: 'q-icons material-icons', fontFamily: 'Material Icons', text: 'directions_car', fontSize: '18px', color: '#fff' }"
		/>
	</GMapMap>
</template>

<script>
export default {
	props: {
		center: String,
	}
}
</script>
