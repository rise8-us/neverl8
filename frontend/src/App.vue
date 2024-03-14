<template>
  <div id="app">
    <CalendarView @dateSelected="handleDateSelected" />
    <ScheduleMeeting v-if="meeting" :meeting="meeting" @scheduleConfirmed="handleScheduleConfirmed" />
  </div>
</template>

<script>
import CalendarView from './components/CalendarView.vue';
import ScheduleMeeting from './components/ScheduleMeeting.vue';
import axios from 'axios';

export default {
  name: 'App',
  components: {
    CalendarView,
    ScheduleMeeting,
  },
  data() {
    return {
      meeting: null,
    };
  },
  methods: {
    async handleDateSelected(date) {
      // Fetch meeting details based on hardcoded ID
      try {
        const response = await axios.get(`/api/meeting?id=26`);
        this.meeting = { ...response.data, selectedDate: date };
      } catch (error) {
        console.error('Error fetching meeting:', error);
      }
    },
    handleScheduleConfirmed() {
      // You could update the meeting object or UI here to reflect the scheduled state
      alert("Meeting scheduled!");
    }
  },
};
</script>
