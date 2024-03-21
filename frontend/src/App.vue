<template>
  <v-app>
    <v-container>
      <v-row class="d-flex justify-center">
        <h1>Neverl8 Meeting Scheduler</h1></v-row
      >
      <v-divider class="my-4"></v-divider>
      <v-row>
        <v-col>
          <CalendarView @dateSelected="handleDateSelected" />
        </v-col>
        <v-col>
          <ScheduleMeeting
            v-if="meeting"
            :meeting="meeting"
            :availableTimeSlots="availableTimeSlots"
            @scheduleMeeting="handleScheduleMeeting"
          /> </v-col
      ></v-row>
    </v-container>
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
    async handleScheduleMeeting(
      selectedTimeSlot,
      candidateName,
      candidateEmail
    ) {
      try {
        await axios.post(`/api/meeting/schedule`, {
          meeting_id: this.meeting.id,
          time_slot: selectedTimeSlot,
          candidate_name: candidateName,
          candidate_email: candidateEmail,
        });
      } catch (error) {
        console.error("Error scheduling meeting:", error);
      }
      alert("Meeting scheduled!");
    },
  },
};
</script>
