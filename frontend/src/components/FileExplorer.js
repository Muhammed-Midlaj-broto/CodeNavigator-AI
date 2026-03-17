import React from "react";

function FileExplorer({ tree, onFileClick }) {

  if (!tree || tree.length === 0) {
    return <div>No files found</div>;
  }

  const renderNode = (node, path = "") => {

    const currentPath = path ? `${path}/${node.name}` : node.name;

    if (node.type === "blob") {
      return (
        <div
          key={currentPath}
          style={{ cursor: "pointer", paddingLeft: "10px" }}
          onClick={() => onFileClick(currentPath)}
        >
          📄 {node.name}
        </div>
      );
    }

    return (
      <div key={currentPath} style={{ paddingLeft: "10px" }}>
        📁 {node.name}

        {node.children &&
          node.children.map((child) =>
            renderNode(child, currentPath)
          )}
      </div>
    );
  };

  return <div>{tree.map((node) => renderNode(node))}</div>;
}

export default FileExplorer;