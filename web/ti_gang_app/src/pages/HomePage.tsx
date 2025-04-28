import { useEffect, useState } from "react";
import { getPunchCount } from "../utils/storage";
import { PunchButton } from "../components/PunchButton";
import { TodayStats } from "../components/TodayStats";

export function HomePage() {
  const [count, setCount] = useState<number>(0);

  useEffect(() => {
    setCount(getPunchCount());
  }, []);

  const refreshCount = () => {
    setCount(getPunchCount());
  };

  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-gray-50">
      <h1 className="text-5xl font-extrabold mb-8 text-gray-800">提肛 APP</h1>
      <TodayStats count={count} />
      <PunchButton onPunch={refreshCount} />
    </div>
  );
}
