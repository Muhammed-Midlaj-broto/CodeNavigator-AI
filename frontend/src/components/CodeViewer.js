import React from "react";

function CodeViewer({ code, onLineClick }) {

  if (!code) {
    return (
      <div style={{ color: "#94a3b8" }}>
        Select a file to view code
      </div>
    );
  }

  const lines = code.split("\n");

  return (

    <div
      style={{
        fontFamily: "monospace",
        fontSize: "14px",
        overflowY: "auto",
        maxHeight: "600px"
      }}
    >

      {lines.map((line, i) => (

        <div
          key={i}
          onClick={() => onLineClick(line, i + 1)}
          style={{
            display: "flex",
            cursor: "pointer",
            padding: "2px 6px"
          }}
        >

          {/* Line number */}
          <span
            style={{
              width: "40px",
              color: "#60a5fa",
              userSelect: "none"
            }}
          >
            {i + 1}.
          </span>

          {/* Code */}
          <span>{line}</span>

        </div>

      ))}

    </div>

  );
}

export default CodeViewer;