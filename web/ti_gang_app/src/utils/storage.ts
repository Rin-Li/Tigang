export const getTodayKey = () => {
    const today = new Date();
    return `${today.getFullYear()}-${today.getMonth() + 1}-${today.getDate()}`;
  };
  
  export const getPunchCount = (): number => {
    const key = getTodayKey();
    const count = localStorage.getItem(key);
    return count ? parseInt(count, 10) : 0;
  };
  
  export const addPunch = () => {
    const key = getTodayKey();
    const count = getPunchCount();
    localStorage.setItem(key, (count + 1).toString());
  };