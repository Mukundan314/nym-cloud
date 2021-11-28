import React from "react";
import * as styles from "./index.module.css";
import image from "../../assets/profile.jpg";

const Profile = () => (
  <div className={styles.profile}>
    <img src={image} alt="profile icon" className={styles.icon} />
  </div>
);

export default Profile;
