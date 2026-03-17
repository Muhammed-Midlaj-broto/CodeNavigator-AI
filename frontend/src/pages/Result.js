import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import FileExplorer from "../components/FileExplorer";
import CodeViewer from "../components/CodeViewer";
import AIExplanation from "../components/AIExplanation";
import LanguageChart from "../components/LanguageChart";
import "./Result.css";

function Result() {
  const navigate = useNavigate();
  const [repoData, setRepoData] = useState(null);
  const [code, setCode] = useState("");
  const [explanation, setExplanation] = useState("");
  const [loading, setLoading] = useState(false);
  const [query, setQuery] = useState("");

  useEffect(() => {
    const stored = localStorage.getItem("repoData");
    if (stored) {
      setRepoData(JSON.parse(stored));
    }
  }, []);

  // Updated "No Data" view with consistent styling
  if (!repoData) {
    return (
      <div className="result-page" style={{ display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center' }}>
        <div className="glass-card" style={{ textAlign: 'center' }}>
          <h2>No repository data found</h2>
          <p>Please start a new analysis to view results.</p>
          <button className="btn-new-search" onClick={() => navigate("/")} style={{ margin: '0 auto' }}>
            Go Back Home
          </button>
        </div>
      </div>
    );
  }

  const openFile = async (path) => {
    try {
      setLoading(true);
      const res = await fetch(
        `http://localhost:8081/explain-file?owner=${repoData.Owner}&repo=${repoData.RepoName}&path=${encodeURIComponent(path)}`
      );
      const data = await res.json();
      setCode(data.code || "");
      setExplanation("Select a line to get detailed explanation.");
    } catch (err) {
      console.log(err);
    } finally {
      setLoading(false);
    }
  };

  const explainLine = async (line, number) => {
    try {
      setLoading(true);
      const res = await fetch("http://localhost:8081/explain-line", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ code: line })
      });
      const data = await res.json();
      setExplanation(`Line ${number}: ${data.explanation}`);
    } catch (err) {
      console.log(err);
    } finally {
      setLoading(false);
    }
  };

  const searchCode = async () => {
    if (!query) return;
    try {
      setLoading(true);
      const res = await fetch("http://localhost:8081/ai-search", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ query: query, code: code })
      });
      const data = await res.json();
      setExplanation(data.result);
    } catch (err) {
      console.log(err);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="result-page">
      <div className="results-container">
        <header className="result-header">
          <h1>
            Results for <span>{repoData.RepoName}</span>
          </h1>
          {/* Changed class to match CSS */}
          <button className="btn-new-search" onClick={() => navigate("/")}>
             <span>+</span> New Search
          </button>
        </header>

        <div className="stats-row">
          <div className="stat-card">
            <small>Total Files</small>
            <h2>{repoData.TotalFiles}</h2>
          </div>
          <div className="stat-card">
            <small>Folders</small>
            <h2>{repoData.TotalFolders}</h2>
          </div>
          <div className="chart-card">
            <LanguageChart languages={repoData.Languages} />
          </div>
        </div>

        <div className="ai-search">
          <input
            type="text"
            placeholder="Ask AI to find code (example: where is login validation)"
            value={query}
            onChange={(e) => setQuery(e.target.value)}
          />
          {/* Added class for consistent button styling */}
          <button className="btn-ai-action" onClick={searchCode} disabled={loading}>
            {loading ? "Searching..." : "AI Search"}
          </button>
        </div>

        <div className="main-grid">
          <div className="glass-panel">
            <h3>Explorer</h3>
            <FileExplorer
              tree={repoData.Tree?.children || []}
              onFileClick={openFile}
            />
          </div>

          <div className="glass-panel">
            <h3>Code Viewer</h3>
            <CodeViewer
              code={code}
              onLineClick={explainLine}
            />
          </div>

          <div className="glass-panel">
            <h3>AI Explanation</h3>
            {loading ? (
              <div className="loader-container">
                <p>AI is thinking...</p>
              </div>
            ) : (
              <AIExplanation explanation={explanation} />
            )}
          </div>
        </div>
      </div>
    </div>
  );
}

export default Result;