import styles from '../styles/identity.module.scss';

const IdentityLayout = ({ children }) => {
  return (
    <>
      <div className="fixed text-3xl font-bold text-white w-screen text-center mt-4">
        LOGO HERE
      </div>
      <div
        className={`h-screen flex flex-col justify-center ${styles.main_background}`}
      >
        {children}
      </div>
    </>
  );
};

export default IdentityLayout;
