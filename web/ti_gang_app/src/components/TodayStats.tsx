interface TodayStatsProps {
    count: number;
  }
  
  export function TodayStats({ count }: TodayStatsProps) {
    return (
      <div className="text-2xl mb-6">
        今天已经提肛 <span className="font-bold">{count}</span> 次！
      </div>
    );
  }
  