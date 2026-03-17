import { useState } from "react";

export default function RepoInput({ setRepoData }) {

  const [url, setUrl] = useState("");

  const analyzeRepo = async () => {

    const res = await fetch("http://localhost:8081/analyze", {
      method: "POST",
      headers: {
        "Content-Type": "application/x-www-form-urlencoded"
      },
      body: new URLSearchParams({
        repo_url: url
      })
    });

    const data = await res.json();

    setRepoData(data);
  };

  return (
    <div>
      <input
        type="text"
        placeholder="GitHub repo URL"
        value={url}
        onChange={(e)=>setUrl(e.target.value)}
      />

      <button onClick={analyzeRepo}>
        Analyze
      </button>
    </div>
  );
}