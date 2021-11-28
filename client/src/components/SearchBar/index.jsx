import React from "react";
import * as styles from "./index.module.css";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faSearch } from "@fortawesome/free-solid-svg-icons";

const SearchBar = () => (
  <div className={styles.searchBar}>
    <div className={styles.icon}>
      <FontAwesomeIcon icon={faSearch} />
    </div>
    <input
      type="text"
      placeholder="Search files and folders"
      className={styles.input}
    />
  </div>
);

export default SearchBar;
