<template>
    <div>
      <input type="date" v-model="selectedDate" @change="fetchMeetingsForDay" />
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        selectedDate: '',
      };
    },
    methods: {
      fetchMeetingsForDay() {
        // Emit an event with the selected date
        this.$emit('dateSelected', this.selectedDate);
      },
    },
  };
  </script>
  