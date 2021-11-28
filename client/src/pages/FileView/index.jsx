import React from "react";
import * as styles from "./index.module.css";
import SearchBar from "../../components/SearchBar";

const FileView = () => (
  <div className={styles.file}>
    <div className={styles.toolBar}>
      <SearchBar />
    </div>
  </div>
);

export default FileView;
