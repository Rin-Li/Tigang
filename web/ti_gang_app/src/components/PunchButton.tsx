import { addPunch } from "../utils/storage";

interface PunchButtonProps {
  onPunch: () => void;
}

export function PunchButton({ onPunch }: PunchButtonProps) {
  const handlePunch = () => {
    addPunch();
    onPunch();
  };

  return (
    <button
      onClick={handlePunch}
      className="bg-blue-500 hover:bg-blue-600 text-white font-bold py-4 px-8 rounded-full text-xl"
    >
      提肛打卡
    </button>
  );
}
