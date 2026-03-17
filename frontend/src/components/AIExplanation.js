import React from "react";

function AIExplanation({ explanation }) {

  if (!explanation) {
    return (
      <div style={{ color: "#94a3b8" }}>
        AI explanation will appear here
      </div>
    );
  }

  const paragraphs = explanation.split("\n");

  return (

    <div
      style={{
        lineHeight: "1.6",
        fontSize: "14px",
        whiteSpace: "pre-wrap"
      }}
    >

      {paragraphs.map((p, i) => (
        <p key={i}>{p}</p>
      ))}

    </div>

  );
}

export default AIExplanation;