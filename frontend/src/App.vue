<template>
  <v-app>
    <CalendarView @dateSelected="handleDateSelected" />
    <ScheduleMeeting
      v-if="meeting"
      :meeting="meeting"
      :availableTimeSlots="availableTimeSlots"
      @scheduleConfirmed="handleScheduleConfirmed"
    />
  </v-app>
</template>

<script>
import CalendarView from "./components/CalendarView.vue";
import ScheduleMeeting from "./components/ScheduleMeeting.vue";
import axios from "axios";

export default {
  name: "App",
  components: {
    CalendarView,
    ScheduleMeeting,
  },
  data() {
    return {
      meeting: null,
      availableTimeSlots: [],
    };
  },

  mounted() {
    this.fetchMeetingDetails();
  },
  methods: {
    async fetchMeetingDetails() {
      // Fetch meeting details based on hardcoded ID
      try {
        const response = await axios.get(`/api/meeting?id=1`);
        this.meeting = response.data;
      } catch (error) {
        console.error("Error fetching meeting:", error);
      }
    },
    async handleDateSelected(date) {
      try {
        const response = await axios.get(
          `/api/meeting/time-slots?date=${date}&id=1`
        );
        this.availableTimeSlots = response.data;
      } catch (error) {
        console.error("Error fetching available time slots:", error);
        this.availableTimeSlots = [];
      }
    },
    handleScheduleConfirmed() {
      alert("Meeting scheduled!");
    },
  },
};
</script>
