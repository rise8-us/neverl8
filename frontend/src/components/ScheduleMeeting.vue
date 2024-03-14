<template>
    <div>
        <h2>Schedule Meeting</h2>
        <div v-if="meeting">
            <p>Meeting: {{ meeting.title }}</p>
            <p>Duration: {{ meeting.duration }} minutes</p>
            <h3>Date Selected: {{ meeting.selectedDate }}</h3>
            <h3>Hosts and Available Slots:</h3>
            <ul>
                <li v-for="host in meeting.Hosts" :key="host.id">
                    {{ host.host_name }}:
                    <ul>
                        <li v-for="slot in calculateTimeSlots(host.time_preferences)" :key="slot.startTime">
                            {{ slot.start }} to {{ slot.end }}
                            <button @click="selectTimePreference(slot)">Select</button>
                        </li>
                    </ul>
                </li>
            </ul>
            <div v-if="selectedTimePreference">
                <h4>Enter your details</h4>
                <form @submit.prevent="scheduleMeeting">
                    <input type="text" v-model="candidateName" placeholder="Your Name" required />
                    <input type="email" v-model="candidateEmail" placeholder="Your Email" required />
                    <button type="submit">Submit</button>
                </form>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    props: {
        meeting: Object,
    },
    data() {
        return {
            candidateName: '',
            candidateEmail: '',
            selectedTimePreference: null,
        };
    },
    methods: {
        selectTimePreference(timePref) {
            // For simplicity, just store the selected time preference
            this.selectedTimePreference = timePref;
        },
        calculateTimeSlots(timePreferences) {
            // Directly map time preferences to a simpler slot structure
            return timePreferences.map(pref => ({
            start: pref.start_window,
            end: pref.end_window,
            }));
        },
        convertToMinutes(time) {
            const [hours, minutes] = time.split(':').map(Number);
            return hours * 60 + minutes;
        },
        convertMinutesToTime(minutes) {
            const hours = Math.floor(minutes / 60);
            const mins = minutes % 60;
            return `${hours.toString().padStart(2, '0')}:${mins.toString().padStart(2, '0')}`;
        },
        scheduleMeeting() {
            // Trigger an event to indicate the meeting has been "scheduled"
            this.$emit('scheduleConfirmed');
            // Reset form
            this.candidateName = '';
            this.candidateEmail = '';
            this.selectedTimePreference = null;
        },
    },
};
</script>
