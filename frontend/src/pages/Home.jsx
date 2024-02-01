import withAuthRedirect from "../providers/withAuthRedirect";
import React, { useEffect, useState } from "react";
import axios from "axios";
import Card from "../components/Card";

const Home = () => {
  const [stats, setStats] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      axios
        .get("https://www.zhefuz.link:8443/api/v1/protected/stats")
        .then(({ data }) => {
          setStats(data);
          console.log(data);
        });
    };

    const interval = setInterval(fetchData, 1000 * 5); //every 5 secs
    return () => clearInterval(interval);
  }, []);
  return (
    <div className="flex flex-col grow">
      <div className="flex  py-28 justify-center  "></div>
      <div className="w-screen flex justify-center ease-in-out duration-500">
        {stats.map((item) => (
          <Card
            hostname={item.hostname}
            disk_usage={item.disk_usage}
            cpu_usage={item.cpu_usage}
            memory_usage={item.memory_usage}
          />
        ))}
      </div>
    </div>
  );
};
export default withAuthRedirect(Home);
