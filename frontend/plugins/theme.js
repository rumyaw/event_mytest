export default defineNuxtPlugin(({ app }) => {
  const theme = useState('theme', () => 'dark')
  
  // Set theme on initial load
  if (process.client) {
    const savedTheme = localStorage.getItem('theme') || 'dark'
    theme.value = savedTheme
    document.documentElement.classList.toggle('dark', savedTheme === 'dark')
    document.documentElement.classList.toggle('light', savedTheme === 'light')
  }
  
  const toggleTheme = () => {
    const newTheme = theme.value === 'dark' ? 'light' : 'dark'
    theme.value = newTheme
    
    if (process.client) {
      localStorage.setItem('theme', newTheme)
      document.documentElement.classList.toggle('dark', newTheme === 'dark')
      document.documentElement.classList.toggle('light', newTheme === 'light')
    }
  }

  return {
    provide: {
      theme,
      toggleTheme
    }
  }
}) 