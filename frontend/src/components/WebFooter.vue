<template>
    <footer>
        <p>Copyright &copy; Karol Pawlak {{ year }} </p>
    </footer>
</template>

<script>
import { computed } from 'vue'

export default {
  setup () {
    const year = computed(() => new Date().getFullYear())
    return { year }
  }
}
</script>
