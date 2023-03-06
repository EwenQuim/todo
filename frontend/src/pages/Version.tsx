export const Version = () => {
  return (
    <div className="flex flex-col items-center justify-center h-screen">
      <h1 className="text-6xl font-bold">Version</h1>
      <p className="text-2xl font-bold">v{process.env.REACT_APP_VERSION}</p>
    </div>
  );
};
