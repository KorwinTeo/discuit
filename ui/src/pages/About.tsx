import StaticPage from '../components/StaticPage';

const About = () => {

  return (
    <StaticPage className="page-about" title="About" noWrap>
      <div className="about-landing">
        <div className="wrap">
          <h1 className="about-heading heading-highlight">
            A social platform by the users, for the users.
          </h1>
          <h2 className="about-subheading">
            Linkforest is a non-profit, open-source community discussion platform. It’s an alternative
            to Reddit.
          </h2>
        </div>
        <div className="squiggly-line"></div>
      </div>
      <div className="about-rest">
        <div className="wrap">
          <div className="about-section about-mission">
            <p>
              Our mission is to build the first large-scale social media platform where the
              interests of the platform are aligned with the interests of the user—a platform, in
              other words, that’s built on principles of ethical design. At the heart of these
              principles is the idea of giving users the freedom to choose their online social
              experience as they would prefer.
            </p>
            <p>
              Social media platforms have hitherto done the opposite and taken away what little
              control the users had, as it served those companies’ own self-interest, which was and
              still remains, to make as much money as possible, without any regard for, indeed to
              the utter detriment of, the well-being of the user.
            </p>
          </div>
          <div className="about-section about-highlights">
            <div className="about-highlight">
              <span className="is-bold">No ads. No tracking.</span>
              There are no ads, no forms of affiliate marketing, and no tracking anywhere on
              Linkforest. And neither your attention, nor your data, is monetized in any way, shape or
              form.
            </div>
            <div className="about-highlight">
              <span className="is-bold">Enshitification-proof.</span>
              {`Linkforest is a non-profit that's funded entirely by its users through donations. The
              lack of a profit motive—and the lack of any shareholders or investors to answer to—is
              essential in keeping this platform completely aligned with the interests and the
              well-being of its users.`}
            </div>
            <div className="about-highlight">
              <span className="is-bold">Giving agency to users.</span>
              Choice over what appears on your feed. Multiple feeds. A plethora of ways to filter
              content. In short, you have complete control over what you see on Linkforest. (Please
              note that Linkforest is a work in progress and that many of these features are yet to be
              built.)
            </div>
            <div className="about-highlight">
              <span className="is-bold">No dark patterns.</span>
              On Linkforest, there are no nagging popups asking you to sign up. You don’t need an
              account to simply view a page. Images, in their highest quality, can be freely
              downloaded. We don’t manipulate you into using our platform more than you desire to.
            </div>
          </div>
        </div>
      </div>
    </StaticPage>
  );
};

export default About;
