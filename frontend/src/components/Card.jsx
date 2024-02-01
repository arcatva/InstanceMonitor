const Card = ({ hostname, cpu_usage, memory_usage, disk_usage }) => {
  return (
    <div className="text-neutral-500 flex bg-white font-extralight font-sans p-8 rounded-3xl  w-2/12 h-40 m-10 shadow-2xl select-none hover:shadow-lg  ease-in-out duration-500 ">
      <div className="w-full flex flex-col font-normal cursor-pointer hover:text-blue-700  ">
        <div className="text-lg mb-6 font-bold">{hostname}</div>
        <div className="text-sm">CPU Usage: {cpu_usage}%</div>
        <div className="text-sm">Memory Usage: {memory_usage}%</div>
        <div className="text-sm ">Disk Usage: {disk_usage}%</div>
      </div>
    </div>
  );
};

export default Card;
