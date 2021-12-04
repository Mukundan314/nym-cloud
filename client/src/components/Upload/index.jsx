import React from "react";
import * as styles from "./index.module.css";

const Upload = () => (
  <div className={styles.upload}>
    <button type="button" className={styles.uploadButton}>
      Upload a file
    </button>
    <input
      type="file"
      name="myfile"
      onChange={(event) => {
        const file = event.target.files[0];
        setTimeout(() => {
          // eslint-disable-next-line no-console
          console.log("Uploaded");
        }, 3000);
      }}
    />
  </div>
);

export default Upload;
