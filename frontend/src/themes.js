// 主题配置
export const themes = {
  light: {
    name: '浅色',
    colors: {
      primary: '#ffffff',
      secondary: '#f5f5f5',
      accent: '#007bff',
      text: '#333333',
      textSecondary: '#778087',
      border: '#e2e2e2',
      headerBg: '#ffffff',
      headerText: '#333333',
      searchBg: 'rgba(0, 0, 0, 0.1)',
      searchText: '#333333',
      searchPlaceholder: 'rgba(0, 0, 0, 0.4)',
      navHover: '#007bff',
      cardBg: '#ffffff',
      cardShadow: '0 2px 3px rgba(0, 0, 0, 0.1)',
      replyCountBg: '#f0f0f0',
      replyCountText: '#999999',
      topicHover: '#007bff',
      sidebarFooterBg: '#f9f9f9',
      subNavBg: '#f5f5f5',      // 子导航背景色
      subNavText: '#333333',    // 子导航文字颜色
      subNavHover: '#007bff',   // 子导航悬停颜色
      subNavBorder: '#e2e2e2'   // 子导航边框颜色
    }
  },
  dark: {
    name: '深色',
    colors: {
      primary: '#1a1a1a',
      secondary: '#242424',
      accent: '#3a86ff',
      text: '#ffffff',
      textSecondary: '#aaaaaa',
      border: '#444444',
      headerBg: '#000000',
      headerText: '#ffffff',
      searchBg: 'rgba(255, 255, 255, 0.2)',
      searchText: '#ffffff',
      searchPlaceholder: 'rgba(255, 255, 255, 0.6)',
      navHover: '#3a86ff',
      cardBg: '#1a1a1a',
      cardShadow: '0 2px 3px rgba(0, 0, 0, 0.3)',
      replyCountBg: '#333333',
      replyCountText: '#aaaaaa',
      topicHover: '#3a86ff',
      sidebarFooterBg: '#242424',
      subNavBg: '#242424',      // 子导航背景色
      subNavText: '#ffffff',    // 子导航文字颜色
      subNavHover: '#3a86ff',   // 子导航悬停颜色
      subNavBorder: '#444444'   // 子导航边框颜色
    }
  },
  v2ex: {
    name: 'V2EX',
    colors: {
      primary: '#ffffff',
      secondary: '#f5f5f5',
      accent: '#e74c3c',
      text: '#333333',
      textSecondary: '#778087',
      border: '#e2e2e2',
      headerBg: '#000000',
      headerText: '#ffffff',
      searchBg: 'rgba(255, 255, 255, 0.2)',
      searchText: '#ffffff',
      searchPlaceholder: 'rgba(255, 255, 255, 0.6)',
      navHover: '#ffffff',
      cardBg: '#ffffff',
      cardShadow: '0 2px 3px rgba(0, 0, 0, 0.1)',
      replyCountBg: '#f0f0f0',
      replyCountText: '#999999',
      topicHover: '#e74c3c',
      sidebarFooterBg: '#f9f9f9',
      subNavBg: '#f5f5f5',      // 子导航背景色
      subNavText: '#333333',    // 子导航文字颜色
      subNavHover: '#e74c3c',   // 子导航悬停颜色
      subNavBorder: '#e2e2e2'   // 子导航边框颜色
    }
  }
};

// 默认主题
export const defaultTheme = 'v2ex';