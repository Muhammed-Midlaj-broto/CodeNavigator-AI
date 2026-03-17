import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import "./Home.css";

function Home() {
  const [repoUrl, setRepoUrl] = useState("");
  const [loading, setLoading] = useState(false); // Added loading state
  const navigate = useNavigate();

  const analyzeRepo = async () => {
    if (!repoUrl) return alert("Please enter a GitHub URL");
    
    setLoading(true); // Start loading animation
    try {
      const res = await fetch("http://localhost:8081/analyze", {
        method: "POST",
        headers: {
          "Content-Type": "application/x-www-form-urlencoded",
        },
        body: new URLSearchParams({
          repo_url: repoUrl,
        }),
      });

      if (!res.ok) throw new Error("Analysis failed");

      const data = await res.json();
      localStorage.setItem("repoData", JSON.stringify(data));
      navigate("/result");
    } catch (err) {
      console.error("Analyze error:", err);
      alert("Error analyzing repository. Check if the backend is running.");
    } finally {
      setLoading(false); // Stop loading regardless of outcome
    }
  };

  return (
    <div className="home-page">

      <div className="home-container">
        <div className="glass-card">
          <div className="brand-icon">
            <img 
              src="/codenavigator-icon.png" /* Updated to your specific filename */
              alt="CodeNavigator AI" 
              className="logo-glow"
            />
          </div>
          
          <h1>
            Codebase <span>Navigator</span>
          </h1>
          <p>The AI-powered guide for your GitHub repositories.</p>

          <div className="input-group">
            <input
              type="text"
              placeholder="Paste GitHub URL (e.g., https://github.com/user/repo)"
              value={repoUrl}
              onChange={(e) => setRepoUrl(e.target.value)}
              disabled={loading} // Disable input while loading
            />
            <button 
              onClick={analyzeRepo} 
              disabled={loading}
              className={loading ? "btn-loading" : ""}
            >
              {loading ? "Analyzing..." : "Start Analysis"}
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}

export default Home;