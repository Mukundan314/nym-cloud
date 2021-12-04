import React from "react";
import * as styles from "./index.module.css";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faSearch } from "@fortawesome/free-solid-svg-icons";
import propTypes from "prop-types";

const SearchBar = ({ search, setSearch }) => (
  <div className={styles.searchBar}>
    <div className={styles.icon}>
      <FontAwesomeIcon icon={faSearch} />
    </div>
    <input
      value={search}
      type="text"
      placeholder="Search files and folders"
      className={styles.input}
      onChange={(e) => setSearch(e.target.value)}
    />
  </div>
);

SearchBar.propTypes = {
  search: propTypes.string.isRequired,
  setSearch: propTypes.func.isRequired,
};

export default SearchBar;
